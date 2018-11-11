package hotelbooking;

import helper.DBHandler;
import helper.RestHandler;
import model.Customer;
import model.Invoice;
import org.json.JSONObject;

import javax.jws.WebService;
import java.util.Date;

@WebService(endpointInterface = "hotelbooking.HotelBooking")
public class HotelBookingImplementation implements HotelBooking {
    private RestHandler rh = new RestHandler();

    /* Web services */

    @Override
    public int orderRoom(String name, String id_card, String email, int room_type_id, String check_in, String check_out) {
        RestHandler rh = new RestHandler();
        DBHandler dh = new DBHandler("root","");

        /* creates customer object and fills it */
        Customer customer = new Customer();
        customer.setIdentity(id_card);
        customer.setEmail(email);
        customer.setName(name);

        int id_cust = dh.getCustomerId(name,email,id_card);
        if(id_cust == -999){
            /* if customer hasn't, register a new entity */
            String cust_register_url = "http://localhost:8060/customer";
            JSONObject customerToSend = new JSONObject();
            customerToSend.put("name",customer.getName());
            customerToSend.put("id",customer.getIdentity());
            customerToSend.put("email",customer.getEmail());
            rh.POSTRequest(cust_register_url,customerToSend);
            /* sets the new id customer, get it from database */
            id_cust = dh.getCustomerId(name,email,id_card);
        }

        /* checks if customer is already registered */
        customer.setId(id_cust);

        JSONObject objectToSend = new JSONObject();
        objectToSend.put("room_id",room_type_id);
        objectToSend.put("customer_id",customer.getId());
        objectToSend.put("in",check_in);
        objectToSend.put("out",check_out);

        String url = "http://localhost:8060/invoice/";
        rh.POSTRequest(url,objectToSend);

        /* calls payment validation */
        int respond;

        /* dummy */
        respond = 1;

        /* update invoice or delete invoice */
        if(respond == 1){

        } else {

        }

        return 0;
    }

    @Override
    public int cancel(int bookingId) {
        int statusCode;
        /* Get Booking Data */
        String url = "http://localhost:8060/invoice/"+bookingId;
        JSONObject invoiceJSON = rh.getRestObject(url,"GET");
        JSONObject dataJSON = invoiceJSON.getJSONObject("data");
        Invoice invoice = new Invoice();
        invoice.setId(bookingId);
        invoice.fillFromJSON(dataJSON);
        invoice.printInvoice();

        /* Check if time of cancellation request > 1 day of booking time itinerary */
        Date now = new Date();
        long difference =  (invoice.getInDate().getTime()-now.getTime())/86400000;
        if(Math.abs(difference) >= 1){
            /* cancellation cancelled */
            statusCode = 0;
        } else {
            /* cancel booking */
            rh.deleteRESTRequest(url);
            statusCode = 1;
        }
        return statusCode;
    }

    @Override
    public int confirmPayment(int booking_id, long price, String payer_name, int payment_type) {
        int status;
        /* Calls Payment Gateway's External API */
        boolean dummy_resp = true;

        /* if payment is confirmed, change internal booking database */
        if(dummy_resp){
            status = 1;
        } else {
            status = 0;
        }
        return status;
    }


}
