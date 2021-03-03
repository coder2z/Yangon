# Yangon

## daw

```go

type Items struct {
    XMLName xml.Name `xml:"items"`
    Text    string   `xml:",chardata"`
    Item    struct {
        Chardata     string `xml:",chardata"`
        Autocomplete string `xml:"autocomplete,attr"`
        Valid        string `xml:"valid,attr"`
        Title        string `xml:"title"`
        Subtitle     string `xml:"subtitle"`
        Icon         string `xml:"icon"`
        Text         struct {
            Text string `xml:",chardata"`
            Type string `xml:"type,attr"`
        } `xml:"text"`
    } `xml:"item"`
}

```
