package extractor

type IOC struct {
	Source string //Source file of the indicator
	Type   string //Type of IOC (e.g., ip, url)
	Value  string //IOC value matched
}
