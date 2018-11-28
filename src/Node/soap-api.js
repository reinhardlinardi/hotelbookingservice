const http = require('http');
const soap = require('soap');
const express = require('express');
const bodyParser = require('body-parser');

var confirmPayment = {
    ConfirmPaymentServiceImplService: {
        ConfirmPaymentServiceImplPort: {
            confirmPayment: function(args) {
                console.log('ACCESSED!');
                console.log(args);
                return {
                    return: 1
                };
            }
        }
    }
};

var xml = require('fs').readFileSync('./myservice.wsdl', 'utf-8');

//http server example
// var server = http.createServer(function(request,response) {
//     response.write('Hello World!');
//     response.end('404: Not Found: ' + request.url);
// });

// server.listen(8000);
// soap.listen(server, '/wsdl', confirmPayment, xml);

//express server example
var app = express();
//body parser middleware are supported (optional)
app.use(bodyParser.text({type: 'text/*'}));
app.listen(8001, function(){
    //Note: /wsdl route will be handled by soap module
    //and all other routes & middleware will continue to work
    soap.listen(app, '/wsdl', confirmPayment, xml);
});