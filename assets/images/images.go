package images

import (
	_ "image/png" // to load png images

	"github.com/hajimehoshi/ebiten/v2"
)

// background
var (
	TitleBackground  *ebiten.Image
	SelectBackground *ebiten.Image
	SkyBackground    *ebiten.Image
	Curtain          *ebiten.Image

	PauseLayer *ebiten.Image
)

// field parts
// - general
var (
	Goal_back     *ebiten.Image
	Goal_front    *ebiten.Image
	MapBackground *ebiten.Image
)

// - prairie  field
var (
	TilePrairie  *ebiten.Image
	Grass1       *ebiten.Image
	Grass2       *ebiten.Image
	Grass3       *ebiten.Image
	MountainNear *ebiten.Image
	MountainFar  *ebiten.Image
	CloudNear    *ebiten.Image
	CloudFar     *ebiten.Image
	RockNormal   *ebiten.Image
)

var (
	// Foods
	Onigiri *ebiten.Image

	// UI parts
	CharaWindow             *ebiten.Image
	CharaSelectButton       *ebiten.Image
	PauseButton             *ebiten.Image
	StartButton             *ebiten.Image
	UpButton                *ebiten.Image
	DownButton              *ebiten.Image
	AttackButton            *ebiten.Image
	SkillButton             *ebiten.Image
	VolumeOnButton          *ebiten.Image
	VolumeOffButton         *ebiten.Image
	StartTitleButton        *ebiten.Image
	ResultConfettiAnimation []*ebiten.Image
	OnigiriAnimation        []*ebiten.Image

	// character standing image
	KuronaStanding     *ebiten.Image
	KomaStanding       *ebiten.Image
	ShishimaruStanding *ebiten.Image

	// character animation
	KuronaAnimation     []*ebiten.Image
	KomaAnimation       []*ebiten.Image
	ShishimaruAnimation []*ebiten.Image

	// character attack image
	AttackKomaFist       *ebiten.Image
	AttackShishimaruFist *ebiten.Image
	AttackScratch        *ebiten.Image

	// character sp effect
	KuronaSpEffect     *ebiten.Image
	KomaSpEffect       *ebiten.Image
	ShishimaruSpEffect *ebiten.Image

	// character skill cut-in image
	KuronaSpBack     *ebiten.Image
	KomaSpBack       *ebiten.Image
	ShishimaruSpBack *ebiten.Image

	// character map icon
	KuronaMapIcon     *ebiten.Image
	KomaMapIcon       *ebiten.Image
	ShishimaruMapIcon *ebiten.Image

	// character state icon
	SpecialReadyIcon *ebiten.Image
	WalkStateIcon    *ebiten.Image
	StaminaEmptyIcon *ebiten.Image
)

// LoadImages loads all public images.
func LoadImages() error {
	err := loadBackground()
	if err != nil {
		return err
	}

	err = loadFieldParts()
	if err != nil {
		return err
	}

	err = loadCharacters()
	if err != nil {
		return err
	}

	err = loadAnimation()
	if err != nil {
		return err
	}

	return nil
}
