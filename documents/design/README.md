# 設計文書

## 起動処理

`ebiten`内部の挙動を詳しく全て把握できている訳ではないので細部が間違ってるかもだが、大体以下のような処理だと思っている。

```mermaid
sequenceDiagram
  participant main as MainGoRoutine
  participant ebiten as EbitenGoRoutine

  create participant game as Game
  main ->> game: インスタンス生成
  create participant sm as SceneManager
  game ->> sm: インスタンス生成
  main ->> ebiten: RunGame(Game)
  note over main, ebiten: ループ処理にGameを登録して開始

  loop 更新と描画
  ebiten ->> game: Update()
  game ->> sm: Update()
  ebiten ->> game: Draw(screen *ebiten.Image)
  game ->> sm: Draw(screen *ebiten.Image)
  alt 画面サイズ変更あり
    ebiten ->> game: Layout(outsideWidth, outsideHeight int)
  end
  end

  main ->> ebiten: 終了処理
  ebiten ->> game: Close()
  game ->> game: Assetsの開放処理
```

## Titleシーンの更新処理

```mermaid
sequenceDiagram
  participant ebiten as EbitenGoRoutine
  participant game as Game
  participant sm as SceneManager
  participant scene as TitleScene
  participant vc as VolumeChecker
  participant ic as InputChecker

  ebiten ->> game: Update()
  game ->> sm: Update()
  sm ->> scene: Update(state *GameState)
  scene ->> vc: Update()
  scene ->> vc: 音量設定の読み取り
  scene ->> scene: 音量設定の更新

  scene ->> ic: Update()
  scene ->> ic: ユーザー入力の読み取り
  alt ユーザーがスタートボタンを押下
    scene -->> sm: GoTo(&SelectScene)
  end
```

* SceneManager: ゲームの各画面を`Scene`と表現しており、それらの画面遷移を制御する仕組み

### `SceneManager`の懸案事項

`Scene`の`Update`メソッドで`GameState`という状態を示す情報を渡しているが、
現状この内部は`SceneManager`のポインタが入っているのみ。

画面遷移するため`Scene`から`SceneManager`の`GoTo`メソッドを呼び出したかったのだと思うが、これでは各`Scene`が「自分の次はXX画面に遷移する」という情報を知っている必要があり、設計的に微妙。画面遷移の順序を自由にするため、`Scene`同士はできるだけ疎結合にしたい。

`Scene`から`Scene`への遷移については、`SceneManager`に任せたい。

例えば、`Scene`作成時にチャンネルを渡しておいて、そのチャンネルに対して通知すれば`SceneManager`が次の画面への遷移を実行してくれる・・という機構にすれば、各所が名前の通りの働きになるのではないだろうか。その方が画面遷移の制御コードが`SceneManager`に集約できるのと、ジワッと画面遷移するような処理も一箇所で実装するだけで済みそう。

### Titleシーンのクラス図

インターフェイスと実装の関係などを図示する。

```mermaid
classDiagram
  class Scene {
    <<interface>>
    Initialize() error
    Update(state *GameState)
    Draw(screen *ebiten.Image)
    StartMusic(isVolumeOn bool)
    StopMusic() error
    IsVolumeOn() bool
  }

  Scene <|-- TitleScene
  SceneManager --> Scene

  TitleScene --> Disc: 音楽再生を移譲
  TitleScene --> sePlayer: SE再生を移譲
  TitleScene --> VolumeChecker: 音量設定の管理を移譲
  TitleScene --> InputChecker: ユーザー入力の扱いを移譲
  TitleScene --> Curtain: 画面遷移時のジワッと描画処理
```

## Stage01シーンの更新処理

```mermaid
sequenceDiagram
  participant ebiten as EbitenGoRoutine
  participant game as Game
  participant sm as SceneManager
  participant scene as Stage01Scene
  participant vc as VolumeChecker
  participant ic as InputChecker

  ebiten ->> game: Update()
  game ->> sm: Update()
  sm ->> scene: Update(state *GameState)
  scene ->> vc: Update()
  scene ->> vc: 音量設定の読み取り
  scene ->> scene: 音量設定の更新

  scene ->> ic: Update()
  scene ->> ic: ユーザー入力の読み取り
  scene ->> scene: 内部のGameState情報を更新
  alt GameStateがrun
    scene ->> scene: run()
  else GameStateがstageClear
    scene ->> scene: ResultEffectを描画
  else GameStateがgameOver
    alt ユーザーがスタート押下
      scene ->> sm: GoTo(&TitleScene)
    end
  end
```

### Stage01シーンのrun()処理

ゲーム進行中のやや複雑な処理をまとめた内部関数`run()`のシーケンス図を以下に示す。

```mermaid
sequenceDiagram
  participant scene as Stage01Scene
  participant player as プレイヤーキャラ
  participant field as PrairieField

  scene ->> player: GetPosition()
  note over scene,player: プレイヤーキャラがゴールに到着したか確認

  scene ->> player: Update()
  scene ->> player: GetScrollVelocity()
  note over scene,player: プレイヤーキャラの移動による変位を取得

  scene ->> field: Update(scroll *view.Vector)
  note over scene,field: プレイヤーキャラの変位を使ってフィールドの各種オブジェクトの描画位置を更新

  scene ->> player: IsAttacked()
  alt プレイヤーキャラが攻撃を繰り出した
    scene ->> field: AttackObstacles(aRect, power): collided, broken
    note over scene,field: 攻撃の当たり判定領域と攻撃力を渡して、攻撃処理を依頼

    alt 攻撃が障害物にあたった
      scene ->> player: ConsumeStaminaByAttack(collided)
      note over scene,player: 攻撃が当たったことによってスタミナ減少
    end

    alt 障害物を壊した
      scene ->> player: AddTension(broken)
      note over scene,player: 障害物の破壊成功でテンション上昇
    end
  end

  scene ->> player: GetRectangle(): pRect
  note over scene,player: プレイヤーキャラの当たり判定領域を取得

  scene ->> field: IsCollidedWithObstacles(pRect): isBlocked
  scene ->> player: BeBlocked(isBlocked)
  note over scene,field: フィールド障害物とプレイヤーキャラの衝突有無から速度低下有無を設定

  scene ->> field: EatFoods(pRect): stamina, tension
  scene ->> player: Eat(stamina, tension)
  note over scene,field: フィールドの食べ物とプレイヤーキャラの衝突有無からスタミナ、テンション更新

  scene ->> player: GetStamina()
  scene ->> player: GetTension()
  scene ->> scene: ゲージなど各種UI要素の表示情報を更新
```

冒頭にステージクリアやゲームオーバーの判定処理があるが、シーケンス図の見やすさを優先するため省略する。

### Stage01シーンの懸案事項

現在は`Stage01Scene`がメッセージウィンドウやゲーム使うボタンなどのインスタンスを直接持っているが、他のステージの画面でもそれら共通UI部品を使うと思うので、部品化したほうが良さそう。

プレイヤーキャラの更新処理の結果として以下をもらえるようにしてはどうかと思う。

* プレイヤーキャラの変位（GetScrollVelocity相当）
* 攻撃の有無とその領域、威力（IsAttacked相当）
* プレイヤーキャラの当たり判定領域（GetRectangle相当）

これらの情報は`Player`の`Update`処理の後には確定可能。そして、次の`Field`の`Update`処理でこれらの情報も渡すことで、フィールドの各種オブジェクトの位置更新処理の際にあわせて衝突の有無、破壊の有無、スタミナやテンション上昇の有無といった情報を全て更新可能（for文での繰り返し処理を一回で済ませられる）ではないかと思う。

さらに、`Field`の`Update`の結果として`Player`に返す情報をまとめて返せば、その情報を使ってさらに`Player`の更新ができる（`UpdateWithFieldConditions`などでどうだろうか）。

### Stage01シーンのクラス図

```mermaid
classDiagram
  class Field {
    <<interface>>
    Initialize(lanes *Lanes, goalX float64)
    Update(scroll *view.Vector)
    DrawFarther(screen *ebiten.Image)
    DrawCloser(screen *ebiten.Image)
    IsCollidedWithObstacles(hr *view.HitRectangle) bool
    EatFoods(hr *view.HitRectangle) (stamina int, tension int)
    AttackObstacles(hr *view.HitRectangle, power float64) (int, int)
  }

  Field <|-- PrairieLane
```

```mermaid
classDiagram
  class Scene {
    <<interface>>
  }

  class Field {
    <<interface>>
  }

  class charaPlayer {
    - stamina
    - tension
  }

  class VelocityController {
    <<interface>>
  }

  Scene <|-- Stage01Scene
  Field <|-- PrairieField

  Stage01Scene --> charaPlayer
  Stage01Scene --> Field
  Stage01Scene --> Gauge: 各種ゲージUI
  Stage01Scene --> Progress: 進捗表示UI
  Stage01Scene --> Disc
  Stage01Scene --> sePlayer
  Stage01Scene --> VolumeChecker
  Stage01Scene --> InputChecker
  Stage01Scene --> Curtain

  charaPlayer --> StateMachine
  charaPlayer --> StepAnimation
  charaPlayer --> VelocityController

  VelocityController <|-- KuronaVc
  VelocityController <|-- KomaVc
  VelocityController <|-- ShishimaruVc
```

### GameStateのステートマシン図

ゲームシーンのStateについて記述する。

```mermaid
stateDiagram-v2
  [*] --> wait

  wait --> readyCall: スタートボタンが押された
  readyCall --> goCall: Readyボイスの再生が終わった
  goCall --> run: Goボイスの再生が終わった

  run --> pause: Pauseボタンが押された
  pause --> run: 再開ボタンが押された

  run --> skillEffect: プレイヤーキャラのスキルが発動した
  skillEffect --> run: スキル発動エフェクトの再生が終わった

  run --> stageClear: プレイヤーキャラがゴールに到着した
  run --> gameOver: ゴールに未到着で時間切れになった

  stageClear --> [*]
  gameOver --> [*]
```

### プレイヤーStateのステートマシン図

`chara.Player`が内部で`StateMachine`実装に移譲して管理しているプレイヤーキャラのStateについて記述する。

まずは現状を書き出してみる。

```mermaid
stateDiagram-v2
  [*] --> Dash

  Dash --> Walk: 1
  Walk --> Dash: 2
  
  Dash --> Ascending: 3
  Dash --> SkillEffect: 4

  SkillAscending --> Dash: 5
  Ascending --> Dash: 5
  Descending --> Dash: 6
  SkillDescending --> Dash: 6

  Walk --> SkillDash: 2
  SkillDash --> SkillWalk: 1
  SkillDash --> Dash: 7

  SkillWalk --> SkillDash: 2
  SkillWalk --> Dash: 7

  SkillWalk --> SkillAscending: 3
  SkillDash --> SkillAscending: 3
  Walk --> Ascending: 3

  SkillWalk --> SkillDescending: 8
  SkillDash --> SkillDescending: 8
  Walk --> Descending: 8
```

#### パターン一覧

1. スタミナ0または障害物に衝突中
2. スタミナ回復または障害物との衝突解消
3. ユーザーが上昇ボタン押下
4. テンションMaxでユーザーがスキルボタン押下
5. 上のレーンに到着
6. 下のレーンに到着
7. テンション0
8. ユーザーが下降ボタンを押下

スキル発動後はエフェクトやボイスの再生などがあるため、やや状態遷移とそのための操作が複雑になっている。現状は、`Scene`の実装が、`Player`の`UpdateSkillEffect`と`FinishSpEffect`をUpdateのたびに呼び出し、返り値がtrueになったらrun状態に戻す、という実装。（書き出してみて分かったが、上昇・下降前後での状態遷移に不具合がある。）

すでにサウンドの再生にはチャンネルの仕組みを使っており、こちらも同様にチャンネルによる連携にしたい。
スキルエフェクトを担う仕組みからチャンネルを渡しておけば実現できそう。

* `StateMachine`内部では`SkillEffect`に遷移する
* 次回以降の状態遷移では、スキル完了フラグをチェック
* チャンネルから完了通知があれば、スキル完了フラグをtrueにする
* スキル完了フラグがtrueになったのを受けて、状態を`SkillDash`に変更

複雑になっている要因として、内部で「前の状態」を保持・更新して次回の状態遷移に用いている点が挙げられる。
もっとシンプルに、上昇や下降が終了したらまずは`Dash`または`SkillDash`に遷移するようにした方がよいのではないだろうか。

これから作る理想の形としては以下。

```mermaid
stateDiagram-v2
  [*] --> Dash

  Dash --> Walk: 1
  Walk --> Dash: 2
  
  Dash --> Ascending: 3
  Walk --> Ascending: 3
  Ascending --> Dash: 5
  Dash --> Descending: 8
  Walk --> Descending: 8
  Descending --> Dash: 6

  Dash --> SkillEffect: 4
  Walk --> SkillEffect: 4

  SkillEffect --> SkillDash: スキルエフェクト完了
  SkillDash --> SkillWalk: 1
  SkillWalk --> SkillDash: スタミナ回復、障害物除去
  
  SkillDash --> SkillAscending: 3
  SkillWalk --> SkillAscending: 3
  SkillAscending --> SkillDash: 5
  SkillDash --> SkillDescending: 8
  SkillWalk --> SkillDescending: 8
  SkillDescending --> SkillDash: 6

  SkillDash --> Dash: テンション0
  SkillWalk --> Walk: テンション0
```
