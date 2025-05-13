package models

type ProcessQueue struct {
    Ready     []*PCB `json:"ready"`
    Running   []*PCB `json:"running"`
    Waiting   []*PCB `json:"waiting"`
    Backup    []*PCB `json:"backup"`
    Suspended []*PCB `json:"suspended"`
}