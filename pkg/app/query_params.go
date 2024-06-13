package app

type QueryParam struct {
	Name    string `schema:"name"`
	Age     int    `schema:"age"`
	Address string `schema:"address"`
}
