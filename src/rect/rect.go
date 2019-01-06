package rect

type IRect interface {
	GetId() int32

	GetMinX() float32
	GetMinY() float32

	GetMidX() float32
	GetMidY() float32

	GetMaxX() float32
	GetMaxY() float32

	GetW() float32
	GetH() float32

	SetW(w float32)
	SetH(h float32)

	SetMidX(cx float32)
	SetMidY(cy float32)
}

type Rect struct {
	id int32

	minx float32
	miny float32

	midx float32
	midy float32

	maxx float32
	maxy float32

	w, h float32
}

func (rect *Rect) GetId() int32 {
	return rect.id
}

func (rect *Rect) GetMinX() float32 {
	return rect.minx
}

func (rect *Rect) GetMinY() float32 {
	return rect.miny
}

func (rect *Rect) GetMidX() float32 {
	return rect.midx
}

func (rect *Rect) GetMidY() float32 {
	return rect.midy
}

func (rect *Rect) GetMaxX() float32 {
	return rect.maxx
}

func (rect *Rect) GetMaxY() float32 {
	return rect.maxy
}

func (rect *Rect) GetW() float32 {
	return rect.w
}

func (rect *Rect) GetH() float32 {
	return rect.h
}

func (rect *Rect) SetW(w float32) {
	rect.w = w
}

func (rect *Rect) SetH(h float32) {
	rect.h = h
}

func (rect *Rect) SetMidX(cx float32) {
	rect.minx = cx - rect.w / 2
	rect.midx = cx
	rect.maxx = cx + rect.w / 2
}

func (rect *Rect) SetMidY(cy float32) {
	rect.miny = cy - rect.h / 2
	rect.midy = cy
	rect.maxy = cy + rect.h / 2
}

//////////////////////////////////////////////////////////
func CreateRect(id int32, w, h, cx, cy float32) *Rect {
	return &Rect{
		id: id,

		minx: cx - w / 2,
		miny: cy - h / 2,

		midx: cx,
		midy: cy,

		maxx: cx + w / 2,
		maxy: cy + h / 2,

		w: w,
		h: h,
	}
}
