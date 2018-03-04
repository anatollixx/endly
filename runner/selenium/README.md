
**Selenium Runner** 

Selenium runner opens a web session to run a various action on web driver or web elements.

| Service Id | Action | Description | Request | Response |
| --- | --- | --- | --- | --- |
| selenium | start | start standalone selenium server | [ServerStartRequest](service_contract.go) | [ServerStartResponse](service_contract.go) |
| selenium | stop | stop standalone selenium server | [ServerStopRequest](service_contract.go) | [ServerStopResponse](service_contract.go) |
| selenium | open | open a new browser with session id for further testing | [OpenSessionRequest](service_contract.go) | [OpenSessionResponse](service_contract.go) |
| selenium | close | close browser session | [CloseSessionRequest](service_contract.go) | [CloseSessionResponse](service_contract.go) |
| selenium | call-driver | call a method on web driver, i.e wb.GET(url)| [WebDriverCallRequest](service_contract.go) | [ServiceCallResponse](service_contract.go) |
| selenium | call-element | call a method on a web element, i.e. we.Click() | [WebElementCallRequest](service_contract.go) | [WebElementCallResponse](service_contract.go) |
| selenium | run | run set of action on a page | [RunRequest](service_contract.go) | [RunResponse](service_contract.go) |

call-driver and call-element actions's method and parameters are proxied to stand along selenium server via [selenium client](http://github.com/tebeka/selenium)


Selenium run request defines sequence of action. In case a selector is not specified, call method is defined on [WebDriver](https://github.com/tebeka/selenium/blob/master/selenium.go#L213), 
otherwise [WebElement](https://github.com/tebeka/selenium/blob/master/selenium.go#L370) defined by selector.

[Wait](./../../repeatable.go)  provides ability to wait either some time amount or for certain condition to take place, with regexp to extract data

```json

{
  "SessionID":"$SeleniumSessionID",
  "Actions": [
    {
      "Calls": [
        {
          "Method": "Get",
          "Parameters": [
            "http://play.golang.org/?simple=1"
          ]
        }
      ]
    },
    {
      "Selector": {
        "Value": "#code"
      },
      "Calls": [
        {
          "Method": "Clear"
        },
        {
          "Method": "SendKeys",
          "Parameters": [
            "$code"
          ]
        }
      ]
    },
    {
      "Selector": {
        "Value": "#run"
      },
      "Calls": [
        {
          "Method": "Click"
        }
      ]
    },
    {
      "Selector": {
        "Value": "#output",
        "Key": "output"
      },
      "Calls": [
        {
           "Method": "Text",
           "Wait": {
                    "Repeat": 5,
                    "SleepTimeMs": 100,
                    "ExitCriteria": "$value"
           }
        }
      ]
    }
  ]
}
```