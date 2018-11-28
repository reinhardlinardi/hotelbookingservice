const soap = require('soap');
const http = require('http');

var url = 'http://localhost:8001/wsdl?wsdl';
var args = {
    booking_id:'1',
    price:'200000',
    payer_name : 'reinhard',
    payment_type : 'ovo'
};
soap.createClient(url, function(err, client) {
    // console.log(client.confirmPayment);
    client.confirmPayment(args, function(err, result) {
        // if(err) throw new Error(err);
        if(err) console.log(err);
        console.log(result);
    });
});