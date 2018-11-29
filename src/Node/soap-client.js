const soap = require('soap');

var url = 'http://localhost:8000/order_room?wsdl';
var args = {
    name:'Roland',
    identity_number:'112233445566',
    email : 'roland@avatar.com',
    room_type_id : 1,
    check_in:'21-12-2018',
    check_out:'22-12-2018',
    coupon_code:'abcdefg'
};

soap.createClient(url, function(err, client) {
    // console.log(client.confirmPayment);
    client.orderRoom(args, function(err, result) {
        // if(err) throw new Error(err);
        if(err) console.log(err);
        console.log(result);
    });
});

// var url = 'http://167.205.35.211:8080/easypay/PaymentService?wsdl';

// var args = {
//     paymentId:'f26d4d94-6aa4-4a95-8ff3-2d2525f3b6ed',
//     lastEventId:0
// };

// soap.createClient(url, function(err, client) {
//     // console.log(client.confirmPayment);
//     client.getPaymentEvents(args, function(err, result) {
//         // if(err) throw new Error(err);
//         if(err) console.log(err);
//         console.log(result);
//         var events = result.return.events;
//         console.log(events);
//         for(key in events){
//             console.log(events[key]);
//             if(events[key].attributes.type == 'FAILURE'){
//                 console.log(events[key].attributes.reason);
//             } else {
//                 if(events[key].attributes.type == 'SUCCESS'){
//                     console.log('SUCCESS');
//                 } else {
//                     if(events[key].attributes.type == 'OPEN_URL'){
//                         console.log('OPEN_URL');
//                     } else {
    
//                     }
//                 }
//             }
//         }
//     });
// });

// var args = {
//     paymentMethodId :'go_pay',
//     amount:10000
// };


// soap.createClient(url, function(err, client) {
//     client.beginPayment(args, function(err, result) {
//         if(err) console.log(err);
//         console.log(result);
//     });
// });