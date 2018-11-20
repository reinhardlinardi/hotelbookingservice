const axios = require('axios');
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
    url:"http://localhost:8080/engine-rest/process-definition/key/dummy-process/start"
}

client.subscribe('search-room', async function({ task, taskService }) {
  const roomType = task.variable.get('room_type');
  const checkIn = task.variable.get('check_in');
  const checkOut = task.variable.get ('check_out');
  const processVariables = new Variables();

  try {
    let response = await axios.get(BASE_URL+'room/?in='+checkIn+'&out='+checkOut+'&type='+roomType+'/');
    availableRooms = response.data;
    if (availableRooms.length !== 0) {
        processVariables.set('found', true);
    } else {
        processVariables.set('found', false);
    }
  } catch(error) {
    console.log(error);
  }
  await taskService.complete(task, processVariables);
});

client.subscribe('create-invoice', async function({ task, taskService }) {
  const roomType = task.variable.get('room_type');
  const customerId = task.variables.get('customer_id');
  const checkIn = task.variables.get('check_in');
  const checkOut = task.variables.get('check_out');
  const processVariables = new Variables();

  try {
    let response = await axios.get(BASE_URL+'room/?in='+checkIn+'&out='+checkOut+'&type='+roomType+'/');
    room = response.data[0];
    let result = await axios.post(BASE_URL+'invoice/', {
      room_id : room,
      customer_id : customerId,
      in : checkIn,
      out : checkOut
    });
  } catch (error) {
    console.log(error);
  }

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

client.subscribe('request-payment-information', async function({ task, taskService }) {
    const booking_id = task.variables.get('booking_id');
    const price = task.variables.get('price');
    const payer_name = task.variables.get('payer_name');
    const payment_type = task.variables.get('payment_type');
    const processVariables = new Variables();

    try {
        // let response = await axios.post(paymentGateway.url,{

        // });
        //Calls Payment Gateway Service
    } catch(error) {
        console.log(error);
    }
    await taskService.complete(task, processVariables);
});

client.subscribe('update-invoice', async function({ task, taskService }) {
    const invoiceId = task.variables.get('invoice_id');
    const processVariables = new Variables();

    try {
        let response = await axios.put(BASE_URL+'invoice/'+invoiceId);
        invoiceDetails = respose.data;
        if (invoiceDetails.paid) {
            processVariables.set('paid', true);
        } else {
            processVariables.set('paid', false);
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