package lowlevel

/*
#include <fmod.h>
*/
import "C"
import "unsafe"

type Geometry struct {
	cptr *C.FMOD_GEOMETRY
}

/*
   'Geometry' API
*/

func (g *Geometry) Release() error {
	res := C.FMOD_Geometry_Release(g.cptr)
	return errs[res]
}

/*
   Polygon manipulation.
*/

func (g *Geometry) AddPolygon(directocclusion, reverbocclusion float64, doublesided bool, numvertices int, vertices Vector, polygonindex int) error {
	cvertices := vertices.toC()
	cpolygonindex := C.int(polygonindex)
	res := C.FMOD_Geometry_AddPolygon(g.cptr, C.float(directocclusion), C.float(reverbocclusion), getBool(doublesided), C.int(numvertices), &cvertices, &cpolygonindex)
	return errs[res]
}

func (g *Geometry) NumPolygons() (int, error) {
	var numpolygons C.int
	res := C.FMOD_Geometry_GetNumPolygons(g.cptr, &numpolygons)
	return int(numpolygons), errs[res]
}

func (g *Geometry) MaxPolygons() (int, int, error) {
	var maxpolygons, maxvertices C.int
	res := C.FMOD_Geometry_GetMaxPolygons(g.cptr, &maxpolygons, &maxvertices)
	return int(maxpolygons), int(maxvertices), errs[res]
}

func (g *Geometry) PolygonNumVertices(index int) (int, error) {
	var numvertices C.int
	res := C.FMOD_Geometry_GetPolygonNumVertices(g.cptr, C.int(index), &numvertices)
	return int(numvertices), errs[res]
}

func (g *Geometry) SetPolygonVertex(index, vertexindex int, vertex Vector) error {
	cvertex := vertex.toC()
	res := C.FMOD_Geometry_SetPolygonVertex(g.cptr, C.int(index), C.int(vertexindex), &cvertex)
	return errs[res]
}

func (g *Geometry) PolygonVertex(index, vertexindex int) (Vector, error) {
	var cvertex C.FMOD_VECTOR
	var vertex Vector
	res := C.FMOD_Geometry_GetPolygonVertex(g.cptr, C.int(index), C.int(vertexindex), &cvertex)
	vertex.fromC(cvertex)
	return vertex, errs[res]
}

func (g *Geometry) SetPolygonAttributes(index int, directocclusion, reverbocclusion float64, doublesided bool) error {
	res := C.FMOD_Geometry_SetPolygonAttributes(g.cptr, C.int(index), C.float(directocclusion), C.float(reverbocclusion), getBool(doublesided))
	return errs[res]
}

func (g *Geometry) PolygonAttributes(index int) (float64, float64, bool, error) {
	var directocclusion, reverbocclusion C.float
	var doublesided C.FMOD_BOOL
	res := C.FMOD_Geometry_GetPolygonAttributes(g.cptr, C.int(index), &directocclusion, &reverbocclusion, &doublesided)
	return float64(directocclusion), float64(reverbocclusion), setBool(doublesided), errs[res]
}

/*
   Object manipulation.
*/

func (g *Geometry) SetActive(active bool) error {
	res := C.FMOD_Geometry_SetActive(g.cptr, getBool(active))
	return errs[res]
}

func (g *Geometry) IsActive() (bool, error) {
	var active C.FMOD_BOOL
	res := C.FMOD_Geometry_GetActive(g.cptr, &active)
	return setBool(active), errs[res]
}

func (g *Geometry) SetRotation(forward, up Vector) error {
	cforward := forward.toC()
	cup := up.toC()
	res := C.FMOD_Geometry_SetRotation(g.cptr, &cforward, &cup)
	return errs[res]
}

func (g *Geometry) Rotation() (Vector, Vector, error) {
	var cforward, cup C.FMOD_VECTOR
	var forward, up Vector
	res := C.FMOD_Geometry_GetRotation(g.cptr, &cforward, &cup)
	forward.fromC(cforward)
	up.fromC(cup)
	return forward, up, errs[res]
}

func (g *Geometry) SetPosition(position Vector) error {
	cposition := position.toC()
	res := C.FMOD_Geometry_SetPosition(g.cptr, &cposition)
	return errs[res]
}

func (g *Geometry) Position() (Vector, error) {
	var cposition C.FMOD_VECTOR
	var position Vector
	res := C.FMOD_Geometry_GetPosition(g.cptr, &cposition)
	position.fromC(cposition)
	return position, errs[res]
}

func (g *Geometry) SetScale(scale Vector) error {
	cscale := scale.toC()
	res := C.FMOD_Geometry_SetScale(g.cptr, &cscale)
	return errs[res]
}

func (g *Geometry) Scale() (Vector, error) {
	var cscale C.FMOD_VECTOR
	var scale Vector
	res := C.FMOD_Geometry_GetScale(g.cptr, &cscale)
	scale.fromC(cscale)
	return scale, errs[res]
}

// NOTE: Not implement yet
func (g *Geometry) Save(data *interface{}, datasize *C.int) error {
	//FMOD_RESULT F_API FMOD_Geometry_Save(FMOD_GEOMETRY *geometry, void *data, int *datasize);
	return ErrNoImpl
}

/*
   Userdata set/get.
*/

func (g *Geometry) SetUserData(userdata interface{}) error {
	data := *(*[]*C.char)(unsafe.Pointer(&userdata))
	res := C.FMOD_Geometry_SetUserData(g.cptr, unsafe.Pointer(&data))
	return errs[res]
}

// Retrieves the user value that that was set by calling the "Geometry.SetUserData" function.
func (g *Geometry) UserData() (interface{}, error) {
	var userdata *interface{}
	cUserdata := unsafe.Pointer(userdata)
	res := C.FMOD_Geometry_GetUserData(g.cptr, &cUserdata)
	return *(*interface{})(cUserdata), errs[res]
}
