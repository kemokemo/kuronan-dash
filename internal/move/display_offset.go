package move

// 変化量
const (
	gradually  = 0.7
	speedy     = 2.0
	moreSpeedy = 4.0
)

// Target Offsets
const (
	walkX      = -50.0
	dashX      = 0.0
	skillWalkX = 25.0
	skillDashX = 50.0
)

type displayOffset struct {
	prevState        State
	currentState     State
	targetOffset     float64
	currentOffset    float64
	currentAllOffset float64
	factor           float64
	polarity         int
}

func (d *displayOffset) Update(s State) {
	if s == d.currentState ||
		s == SkillEffect || s == Wait || s == Pause {
		return
	}

	d.prevState = d.currentState
	d.currentState = s
	d.updateFactors()
	d.currentOffset = d.factor * float64(d.polarity)
}

func (d *displayOffset) updateFactors() {
	switch d.currentState {
	case Walk:
		d.targetOffset = walkX
		switch d.prevState {
		case Dash, Ascending, Descending:
			d.factor = gradually
		case SkillWalk:
			d.factor = speedy
		case SkillDash:
			d.factor = moreSpeedy
		default:
			d.factor = 0
		}

	case Dash, Ascending, Descending:
		d.targetOffset = dashX
		switch d.prevState {
		case Walk:
			d.factor = gradually
		case SkillWalk:
			d.factor = gradually
		case SkillDash:
			d.factor = speedy
		default:
			d.factor = 0
		}

	case SkillWalk:
		d.targetOffset = skillWalkX
		switch d.prevState {
		case Walk:
			d.factor = speedy
		case Dash, Ascending, Descending:
			d.factor = gradually
		case SkillDash:
			d.factor = gradually
		default:
			d.factor = 0
		}

	case SkillDash, SkillAscending, SkillDescending:
		d.targetOffset = skillDashX
		switch d.prevState {
		case Walk:
			d.factor = moreSpeedy
		case Dash, Ascending, Descending:
			d.factor = speedy
		case SkillWalk:
			d.factor = gradually
		default:
			d.factor = 0
		}
	}

	if d.currentAllOffset > d.targetOffset {
		d.polarity = -1
	} else {
		d.polarity = 1
	}
}

// 毎ターン、キャラクターの表示位置を動かす差分だけ返す。
// 累計が、現在のステータスでの目的位置にたどり着くだけ返した場合はゼロを返す。
func (d *displayOffset) GetXAxisOffset() float64 {
	if d.polarity > 0 && d.currentAllOffset >= d.targetOffset {
		d.currentAllOffset = d.targetOffset
		return 0.0
	}

	if d.polarity < 0 && d.currentAllOffset <= d.targetOffset {
		d.currentAllOffset = d.targetOffset
		return 0.0
	}

	d.currentAllOffset += d.currentOffset
	return d.currentOffset
}
