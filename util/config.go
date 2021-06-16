package util


type MiniSpider struct {
	UrlListFile     string `yaml:"urlListFile"`
	OutputDirectory string `yaml:"outputDirectory"`
	MaxDepth        int    `yaml:"maxDepth"`
	CrawlInterval   int    `yaml:"crawlInterval"`
	CrawlTimeout    int    `yaml:"crawlTimeout"`
	ThreadCount     int    `yaml:"threadCount"`
}