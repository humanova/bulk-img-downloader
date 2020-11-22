# bulk img downloader

simple and efficient image downloader (and pdf generator) with concurrency

![Sample gif](media/sample.gif) *downlading 224 images and generating a pdf (gif in real-time)*


Sample config:

```json
{
  "url" : "https://url/to/file/%s.png",
  "startIdx" : 1,
  "stopIdx" : 10,
  "extension" : "png",
  "makepdf" : true,
  "generateuuid" : false
}
```