const soap = require('soap');

// var url = 'http://localhost:8000/order_room?wsdl';
// var args = {
//     name:'Roland',
//     identity_number:'112233445566',
//     email : 'roland@avatar.com',
//     room_type_id : 1,
//     check_in:'01-12-2018',
//     check_out:'02-12-2018',
//     coupon_code:'abcdefg'
// };

// soap.createClient(url, function(err, client) {
//     // console.log(client.confirmPayment);
//     client.orderRoom(args, function(err, result) {
//         // if(err) throw new Error(err);
//         if(err) console.log(err);
//         console.log(result);
//     });
// });

var url = 'http://167.205.35.211:8080/easypay/PaymentService?wsdl';
var args = {
    paymentId:'d94f85fe-ac80-4e0e-bdab-8947001c2cff',
    lastEventId:68
};

soap.createClient(url, function(err, client) {
    // console.log(client.confirmPayment);
    client.getPaymentEvents(args, function(err, result) {
        // if(err) throw new Error(err);
        if(err) console.log(err);
        console.log(result);
        console.log(result.return.events);
    });
});