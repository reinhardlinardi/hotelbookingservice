const soap = require('soap');
const express = require('express');
const bodyParser = require('body-parser');
const axios = require('axios');

const order_url = 'http://localhost:8080/engine-rest/process-definition/key/order-room/start';
const user_id_url = 'http://167.205.35.220/api/entity/customer';

var orderRoom = {
    OrderRoomServiceImplService: {
        OrderRoomServiceImplPort: {
            orderRoom: async function(args) {
                console.log('ACCESSED!');
                console.log(args);
                //Calls Order room camunda

                const roomTypes = {
                    '1' : 'Single',
                    '2' : 'Double',
                    '3' : 'Family'
                }
                try{
                    let userResult = await axios.get(user_id_url+'?name='+args.name+'&id='+args.identity_number+'&email='+args.email);

                    console.log(userResult);
                    if(userResult.data.success){
                        let result = await axios.post(order_url,{
                            "variables":{
                                "user_id":{
                                    "value" : userResult.data.data.id,
                                    "type" : "integer"
                                },
                                "room_type":{
                                    "value" : roomTypes[args.room_type_id],
                                    "type" : "string"
                                },
                                "check_in":{
                                    "value" : args.check_in,
                                    "type" : "string"
                                },
                                "check_out":{
                                    "value" : args.check_out,
                                    "type" : "string"
                                }
                            }
                        });

                        if(result.data.success){
                            console.log(result.data.data);
                        } else {
                            console.log(result.data.message);
                        }
                    } else {
                        console.log(userResult.data.message);
                    }
                } catch (error){
                    console.log(error);
                }
                return {
                    return: 1
                };
            }
        }
    }
};

var xml = require('fs').readFileSync('../../wsdl/order_room.wsdl', 'utf-8');

//express server example
var app = express();
//body parser middleware are supported (optional)
app.use(bodyParser.text({type: 'text/*'}));
app.listen(8000, function(){
    //Note: /wsdl route will be handled by soap module
    //and all other routes & middleware will continue to work
    soap.listen(app, '/order_room', orderRoom, xml);
});