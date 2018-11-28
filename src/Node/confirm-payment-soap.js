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

var xml = require('fs').readFileSync('../../wsdl/confirm_payment.wsdl', 'utf-8');

//express server example
var app = express();
//body parser middleware are supported (optional)
app.use(bodyParser.text({type: 'text/*'}));
app.listen(4002, function(){
    //Note: /wsdl route will be handled by soap module
    //and all other routes & middleware will continue to work
    soap.listen(app, '/confirm_payment', confirmPayment, xml);
});