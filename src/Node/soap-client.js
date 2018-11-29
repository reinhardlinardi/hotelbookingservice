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
    paymentId:'743bbd9f-287c-466f-9fe8-2d95b987dba2',
    lastEventId:0
};

soap.createClient(url, function(err, client) {
    // console.log(client.confirmPayment);
    client.getPaymentEvents(args, function(err, result) {
        // if(err) throw new Error(err);
        if(err) console.log(err);
        console.log(result);
        var events = result.return.events;
        console.log(events);
        for(key in events){
            console.log(events[key]);
            if(events[key].attributes.type =='FAILURE'){
                console.log(events[key].attributes.reason);
            } else {
                if(events[key].attributes.type == 'SUCCESS'){
                    console.log('SUCCESS');
                } else {
                    if(events[key].attributes.type == 'OPEN_URL'){
                        console.log('SUCCESS');
                    } else {
    
                    }
                }
            }
        }
    });
});

// var args = {
//     paymentMethodId :'ovo',
//     amount:10000
// };


// soap.createClient(url, function(err, client) {
//     client.beginPayment(args, function(err, result) {
//         if(err) console.log(err);
//         console.log(result);
//     });
// });