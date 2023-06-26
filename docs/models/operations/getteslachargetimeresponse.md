# GetTeslaChargeTimeResponse


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ChargeTime`                                                                       | [*shared.ChargeTime](../../models/shared/chargetime.md)                            | :heavy_minus_sign:                                                                 | returns the date and time the vehicle expects to "complete" this charging session. |
| `ContentType`                                                                      | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `StatusCode`                                                                       | *int*                                                                              | :heavy_check_mark:                                                                 | N/A                                                                                |
| `RawResponse`                                                                      | [*http.Response](https://pkg.go.dev/net/http#Response)                             | :heavy_minus_sign:                                                                 | N/A                                                                                |