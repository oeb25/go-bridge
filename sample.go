package main

type Revision struct {
	Value   string `json:"value"`
	Created int64  `json:"created"`
}

type Item struct {
	ID          *int                  `json:"id"`
	Title       string                `json:"title"`
	Validated   bool                  `json:"validated,omitempty"`
	ValidatedAt int64                 `json:"validatedAt"`
	SupplierID  int                   `json:"supplierID"`
	Properties  map[string][]Revision `json:"properties"`
}

type Order struct {
	ID          int          `json:"id"`
	Items       map[int]Item `json:"items"`
	hiddenField int
}
