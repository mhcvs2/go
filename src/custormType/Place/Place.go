package Place

import (
	"fmt"
)

type Place struct {
	latitude, longtitude float64
	Name string
}

func New(latitude, longtitude float64, name string) *Place {
	return &Place{saneAngle(0, latitude), saneAngle(0, longtitude), name}
}

func (place *Place) Latitude() float64 {
	return place.latitude
}

func (place *Place) SetLatitude(latitude float64) {
	place.latitude = saneAngle(place.latitude, latitude)
}

func (place *Place) Longtitude() float64 {
	return place.longtitude
}

func (place *Place) SetLongtitude(longtitude float64) {
	place.longtitude = saneAngle(place.longtitude, longtitude)
}

func (place *Place) String() string {
	return fmt.Sprintf("(%.3f度, %.3f度 ) %q", place.latitude, place.longtitude, place.Name)
}

func (original *Place) Copy() *Place {
	return &Place{original.latitude, original.longtitude, original.Name }
}

func saneAngle(old, new float64) float64 {
	//验证
	return new
}


