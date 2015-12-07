package controllers

import (
    // "encoding/json"
    "strconv"
    "github.com/revel/revel"
)

// User Controller
type User struct {
    RestController
}

// GetAll model data
func (c User) GetAll() revel.Result {
    var resp Response

    // fake multiple data
    var data []Data

    for i:=1; i<3; i++ {

        strI := strconv.Itoa(i)

        // soon, replace with metadata or model rows/column

        _attr := map[string]string{
            "username": "user" + strI,
            "fullname": "User " + strI,
        }

        // soon, replace with joined data or relationsel
        rel := []interface{}{}

        _data := Data{
            Id: i,
            Type: "user",
            Attributes: _attr,
            Relationships: rel,
        }
        data = append(data, _data)
    }

    resp.Data = data
    resp.Included = []interface{}{}

    return c.RenderJson(resp)
}

// GetSingle model data
func (c User) GetSingle(id int) revel.Result {
    var resp Response

    // fake multiple data
    var data Data

    strI := strconv.Itoa(id)

    // soon, replace with metadata or model rows/column
    _attr := map[string]string{
        "username": "user" + strI,
        "fullname": "User " + strI,
    }

    // soon, replace with joined data or relationsel
    rel := []interface{}{}
    data = Data{
        Id: id,
        Type: "user",
        Attributes: _attr,
        Relationships: rel,
    }

    resp.Data = data
    resp.Included = []interface{}{}

    return c.RenderJson(resp)
}