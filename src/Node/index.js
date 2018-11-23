const axios = require('axios').create();
const http = require('http');

const { Client, logger, Variables } = require('camunda-external-task-client-js');

// configuration for the Client:
//  - 'baseUrl': url to the Process Engine
//  - 'logger': utility to automatically log important events
const config = { baseUrl: 'http://localhost:8080/engine-rest', use: logger };

// create a Client instance with custom configuration
const client = new Client(config);

const BASE_URL = "http://localhost:8060/"

const paymentGateway = {
    url:"http://localhost:8080/engine-rest/process-definition/key/validate-payment-task/start"
}

client.subscribe('search-room', async function({ task, taskService }) {
  const roomType = task.variables.get('room_type');
  const checkIn = task.variables.get('check_in');
  const checkOut = task.variables.get ('check_out');
  const processVariables = new Variables();

  try {
    let response = await axios.get(BASE_URL+'room', {
        params: {
          in: checkIn,
          out: checkOut,
          type:roomType
        }
    });
    room = response.data.data;

    if (room.length > 0) {
        processVariables.set('found', 'true');
        processVariables.set('room_id',room[0]);
    } else {
        processVariables.set('found', 'false');
    }
  } catch(error) {
    console.log(error);
  }
  await taskService.complete(task, processVariables);
});

client.subscribe('create-invoice', async function({ task, taskService }) {
  const roomType = task.variables.get('room_type');
  const customerId = task.variables.get('user_id');
  const checkIn = task.variables.get('check_in');
  const checkOut = task.variables.get('check_out');
  const roomId = task.variables.get('room_id');

  const processVariables = new Variables();

    try {
        let result = await axios.post(BASE_URL+'invoice/', {
            room_id : roomId,
            customer_id : customerId,
            in : checkIn,
            out : checkOut
          });
          if(result.data.success){
              if(result.data.data.id !== null){
                  processVariables.set('booking_id',result.data.data.id);
              } else {
                  console.log(result.data.message);
              }
          } else {
              console.log(result.data.message);
          }
    } catch (error) {
        console.log(error);
    }
  await taskService.complete(task, processVariables);
});

client.subscribe('get-room-price', async function({ task, taskService }) {
    const room_id = task.variables.get('room_id');
    const processVariables = new Variables();

    processVariables.set('payment_type','credit');
    //Get Room Price
    try{
        let roomResponse = await axios.get(BASE_URL+'room/'+room_id);
        if(roomResponse.data.success){
            price = roomResponse.data.data.price;
            processVariables.set('price',price)
        } else {
            console.log(roomResponse.data.message);
        }
    } catch(error){
        console.log(error);
    }
    await taskService.complete(task, processVariables);
});

client.subscribe('get-payer-name', async function({ task, taskService }) {
    const customer_id = task.variables.get('user_id');
    const processVariables = new Variables();
    //Get Payer Name
    try{
        let custResponse = await axios.get(BASE_URL+'customer/'+customer_id);
        if(custResponse.data.success){
            if(custResponse.data.data !== null){
                payer_name = custResponse.data.data.name;
                processVariables.set('payer_name',payer_name);
            } else {
                console.log('returns 0 data');
            }
        } else {
            console.log(custResponse.data.message);
        }
    } catch(error){
        console.log(error);
    }
    await taskService.complete(task, processVariables);
});

client.subscribe('cancel-booking', async function({ task, taskService }) {
    const invoiceId = task.variables.get('invoice_id');
    const processVariables = new Variables();

    try {
        let response = await axios.delete(BASE_URL+'invoice/'+invoiceId);
        bookingDetails = response.data;
        if (bookingDetails.length !== 0) {
            //Sets Payment status
            if(bookingDetails.paid === true){
                processVariables.set('paid',true);
            } else {
                processVariables.set('paid',false);
            }

            //Calls cancel API
            cancelResponse = await axios.delete(BASE_URL+'invoice/'+invoiceId);
            if(cancelResponse.data.success === true){
                processVariables.set('cancelled', true);
            } else {
                console.log(cancelResponse.data.message);
                processVariables.set('cancelled',false);
            }
        } else {
            processVariables.set('cancelled', false);
        }
    } catch(error) {
        console.log(error);
    }
    await taskService.complete(task, processVariables);
});

client.subscribe('request-payment-confirmation', async function({ task, taskService }) {
    const booking_id = task.variables.get('booking_id');
    const price = task.variables.get('price');
    const payer_name = task.variables.get('payer_name');
    const payment_type = task.variables.get('payment_type');
    const processVariables = new Variables();

    try {
        //Calls Payment Gateway Service
        processVariables.set('approved',true);
    } catch(error) {
        console.log(error);
    }
    await taskService.complete(task, processVariables);
});

client.subscribe('update-invoice', async function({ task, taskService }) {
    var invoiceId = task.variables.get('booking_id');
    const processVariables = new Variables();
    console.log(typeof invoiceId);
    if (!(typeof invoiceId === 'string' || invoiceId instanceof String)){
        invoiceId = Number(invoiceId)
    }

    try {
        let response = await axios.put(BASE_URL+'invoice/'+invoiceId);
        invoiceDetails = response.data;
        if (invoiceDetails.success) {
            processVariables.set('invoice-updated', true);
        } else {
            console.log(invoiceDetails.message);
            processVariables.set('invoice-updated', false);
        }
    } catch(error) {
        console.log(error);
    }
    await taskService.complete(task, processVariables);
});

client.subscribe('dummy-validation', async function({ task, taskService }) {
    const price = task.variables.get('price');
    const name = task.variables.get('name');
    const type = task.variables.get('type');
    const confirmation = task.variables.get('confirm');
    const processVariables = new Variables();

    var verdict = confirmation;

    if(name === "reinhard"){
        verdict = true;
    } else if(price>5000000){
        verdict = true;
    }
    console.log(verdict);

    if(verdict){
        processVariables.set('confirm',true);
    } else {
        processVariables.set('confirm',false);
    }
    await taskService.complete(task, processVariables);
});