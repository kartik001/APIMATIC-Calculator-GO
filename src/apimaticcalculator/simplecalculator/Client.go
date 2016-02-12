/*
 * apimaticcalculator
 *
 * This file was automatically generated by APIMATIC BETA v2.0 on 02/12/2016
 */
package simplecalculator

import(
    "strconv"
    "github.com/apimatic/unirest-go"
    "apimaticcalculator"
    "apimaticcalculator/apihelper"
    "apimaticcalculator/models"
)

/*
 * Input structure for the method GetCalculate
 */
type GetCalculateInput struct {
    Operation       models.OperationTypeEnum //The operator to apply on the variables
    X               float32         //The LHS value
    Y               float32         //The RHS value
}

/*
 * Client structure as interface implementation
 */
type SIMPLECALCULATOR_IMPL struct { }

/**
 * Calculates the expression using the specified operation.
 * @param  GetCalculateInput     Structure with all inputs
 * @return	Returns the float32 response from the API call
 */
func (me *SIMPLECALCULATOR_IMPL) GetCalculate (input *GetCalculateInput) (float32, error) {
    //the base uri for api requests
    queryBuilder := apimaticcalculator.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/{operation}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithTemplateParameters(queryBuilder, map[string]interface{} {
        "operation" : models.OperationTypeEnumToValue(input.Operation),
    }) 
    if err != nil {
        //error in template param handling
        return 0.0, err
    }

    
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithQueryParameters(queryBuilder, map[string]interface{} {
        "x" : input.X,
        "y" : input.Y,
    })
    if err != nil {
        //error in query param handling
        return 0.0, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return 0.0, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return 0.0, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return 0.0, err
    }
    
    //returning the response
    var retVal float32
    retVal, err = strconv.ParseFloat(response.Body, 32)
    if err != nil {
        //error in parsing
        return 0.0, err
    }
    return retVal, nil
}

