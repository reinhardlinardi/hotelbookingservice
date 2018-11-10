package hotelbooking;

import helper.RestHandler;
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
        /* Calls Payment Gateway's External API */

        /* if payment is confirmed, change internal booking database */
        return 0;
    }


}
