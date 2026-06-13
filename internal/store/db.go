package store

type Cluster struct {
    ID   int
    Name string
}

func NewCluster(id int, name string) Cluster {
    return Cluster{ID: id, Name: name}
}
