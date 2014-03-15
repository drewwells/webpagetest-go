package wpt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type WPTResult struct {
	Run                        int32              `json:"run"`
	URL                        string             `json:"URL"`
	LoadTime                   int32              `json:"loadTime"`
	TTFB                       int32              `json:"TTFB"`
	BytesOut                   int32              `json:"bytesOut"`
	BytesOutDoc                int32              `json:"bytesOutDoc"`
	Connections                int32              `json:"connections"`
	Requests                   []interface{}      `json:"requests"`
	RequestsDoc                int32              `json:"requestsDoc"`
	Responses_200              int32              `json:"responses_200"`
	Responses_404              int32              `json:"responses_404"`
	Responses_other            int32              `json:"responses_other"`
	Result                     int32              `json:"result"`
	Render                     int32              `json:"render"`
	FullyLoaded                int32              `json:"fullyLoaded"`
	Cached                     int32              `json:"cached"`
	DocTime                    int32              `json:"docTime"`
	DomTime                    int32              `json:"domTime"`
	Score_cache                int32              `json:"score_cache"`
	Score_cdn                  int32              `json:"score_cdn"`
	Score_gzip                 int32              `json:"score_gzip"`
	Score_cookies              int32              `json:"score_cookies"`
	Score_keep_alive           int32              `json:"score_keep-alive"`
	Score_minify               int32              `json:"score_minify"`
	Score_combine              int32              `json:"score_combine"`
	Score_compress             int32              `json:"score_compress"`
	Score_etags                int32              `json:"score_etags"`
	Gzip_total                 float64            `json:"gzip_total"`
	Gzip_savings               int32              `json:"gzip_savings"`
	Minify_total               int32              `json:"minify_total"`
	Minify_savings             int32              `json:"minify_savings"`
	Image_total                int32              `json:"image_total"`
	Image_savings              int32              `json:"image_savings"`
	Optimization_checked       int32              `json:"optimization_checked"`
	Aft                        int32              `json:"aft"`
	DomElements                int32              `json:"domElements"`
	PageSpeedVersion           float32            `json:"pageSpeedVersion,string"`
	Title                      string             `json:"title"`
	TitleTime                  int32              `json:"titleTime"`
	LoadEventStart             int32              `json:"loadEventStart"`
	LoadEventEnd               int32              `json:"loadEventEnd"`
	DomContentLoadedEventStart int32              `json:"domContentLoadedEventStart"`
	DomContentLoadedEventEnd   int32              `json:"domContentLoadedEventEnd"`
	LastVisualChange           int32              `json:"lastVisualChange"`
	Browser_name               string             `json:"browser_name"`
	Browser_version            string             `json:"browser_version"`
	Server_count               int32              `json:"server_count"`
	Server_rtt                 int32              `json:"server_rtt"`
	Base_page_cdn              string             `json:"base_page_cdn"`
	Adult_site                 int32              `json:"adult_site"`
	Fixed_viewport             int32              `json:"fixed_viewport"`
	Score_progressive_jpeg     int32              `json:"score_progressive_jpeg"`
	FirstPaint                 int32              `json:"firstPaint"`
	DocCPUms                   float32            `json:"docCPUms"`
	FullyLoadedCPUms           float32            `json:"fullyLoadedCPUms"`
	DocCPUpct                  float32            `json:"docCPUpct"`
	FullyLoadedCPUpct          float32            `json:"fullyLoadedCPUpct"`
	Date                       float64            `json:"date"`
	SpeedIndex                 int32              `json:"SpeedIndex"`
	VisualComplete             int32              `json:"visualComplete"`
	userTime                   map[string]float32 `json:"userTime"`
}

type Pages struct {
	Details    string `json:"details"`
	Checklist  string `json:"checklist"`
	Breakdown  string `json:"breakdown"`
	Domains    string `json:"domains"`
	ScreenShot string `json:"screenShot"`
}

type Thumbnails struct {
	WaterFall  string `json:"waterfall"`
	Checklist  string `json:"checklist"`
	ScreenShot string `json:"screenShot"`
}

type Images struct {
	Thumbnails
	ConnectionView string `json:"connectionView"`
}

type RawData struct {
	Headers      string `json:"headers"`
	PageData     string `json:"pageData"`
	RequestsData string `json:"requestsData"`
	Utilization  string `json:"utilization"`
}

type VideoFrame struct {
	Time             int32  `json:"time"`
	Image            string `json:"image"`
	VisuallyComplete int32
}

type WPTResultSet struct {
	WPTResult
	Pages       Pages        `json:"pages"`
	Thumbnails  Thumbnails   `json:"thumbnails"`
	Images      Images       `json:"images"`
	RawData     RawData      `json:"rawData"`
	VideoFrames []VideoFrame `json:"videoFrames"`
}

type Views struct {
	FirstView  WPTResult //`json:"firstView"`
	RepeatView WPTResult //`json:"repeatView"`
}

type WPTRun struct {
	FirstView WPTResultSet `json:"firstView"`
	//RepeatView WPTResultSet `json:"repeatView"`
	Id int32 `json:"id"`
}

type WPTBaseResultData struct {
	StatusCode       int32
	StatusText       string
	TestId           string `json:"testId"`
	Summary          string `json:"summary"`
	Location         string `json:"location"`
	Connectivity     string `json:"connectivity"`
	BwDown           int32  `json:"bwDown"`
	BwUp             int32  `json:"bwUp"`
	Latency          int32  `json:"latency"`
	Plr              string //`json:"plr"`
	Completed        float64
	SuccessfulFVRuns int32 `json:"successfulFVRuns"`
	//Average           Views  `json:"average"`
	//Median            Views  `json:"median"`
	//StandardDeviation Views  `json:"standardDeviation"`
}

type WPTResultData struct {
	WPTBaseResultData
	StatusCode int32
	StatusText string
	Runs       map[string]WPTRun `json:"runs"`
}

type WPTResultCleanData struct {
	WPTBaseResultData
	Runs []WPTRun
}

type ResultJSON struct {
	StatusCode int32         `json:"statusCode"`
	StatusText string        `json:"statusText"`
	Completed  float64       `json:"completed"`
	Data       WPTResultData `json:"data"`
}

type Result struct {
	ResultJSON
	Data       WPTResultCleanData
	StatusCode int32
	StatusText string
}

func ProcessResult(response []byte, err error) Result {
	var jsonR ResultJSON
	var result Result

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(response, &jsonR)

	if err != nil {
		log.Fatal(err)
	}

	//Lots of work to unfuck {"0":{},"1":{}} to [{},{}]
	result.StatusCode = jsonR.Data.StatusCode
	result.StatusText = jsonR.Data.StatusText
	result.Data.TestId = jsonR.Data.TestId
	result.Data.Summary = jsonR.Data.Summary
	result.Data.Location = jsonR.Data.Location
	result.Data.Connectivity = jsonR.Data.Connectivity
	result.Data.BwDown = jsonR.Data.BwDown
	result.Data.BwUp = jsonR.Data.BwUp
	result.Data.Latency = jsonR.Data.Latency
	result.Data.Plr = jsonR.Data.Plr
	result.Data.Completed = jsonR.Data.Completed
	result.Data.SuccessfulFVRuns = jsonR.Data.SuccessfulFVRuns

	for _, val := range jsonR.Data.Runs {
		result.Data.Runs = append(result.Data.Runs, val)
	}

	return result
}

func GetResult(url string, key string) (result Result) {

	res, err := http.Get(url + "/jsonResult.php?test=" + key)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return ProcessResult(ioutil.ReadAll(res.Body))
}
