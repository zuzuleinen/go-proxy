package application

type Database struct {
	rows map[string]Row
}

type Row struct {
	Request  []byte
	Response []byte
}

var d *Database

func DB() *Database {
	if d == nil {
		d = &Database{}
		d.rows = make(map[string]Row)
		return d
	}
	return d
}

func (d *Database) Save(id string, req, resp []byte) {
	d.rows[id] = Row{
		Request:  req,
		Response: resp,
	}
}

func (d *Database) Rows() map[string]Row {
	return d.rows
}
