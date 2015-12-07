package controllers

import (
    "github.com/revel/revel"
)

type RestController struct {
    *revel.Controller
}

type ErrorMessage struct {
    Code            int             `json:"code"`
    Details         string          `json:"details"`
}

type Response struct {
    Data            interface{}     `json:"data"`
    Included        interface{}     `json:"included"`
}

type Data struct {
    Id              int             `json:"id"`
    Type            string          `json:"type"`
    Attributes      interface{}     `json:"attributes"`
    Relationships   interface{}     `json:"relationships"`
}