const soap = require('soap');

var url = 'http://localhost:8000/order_room?wsdl';
var args = {
    name:'Roland',
    identity_number:'112233445566',
    email : 'roland@avatar.com',
    room_type_id : 1,
    check_in:'01-12-2018',
    check_out:'02-12-2018',
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