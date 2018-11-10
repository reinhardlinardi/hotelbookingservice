package hotelbooking;

import helper.RestHandler;
import model.Invoice;
import org.json.JSONObject;

import javax.jws.WebService;

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
        /* Get Booking Data */
        String getInvoiceURL = "http://localhost:8060/invoice/"+bookingId;
        JSONObject invoiceJSON = rh.getRestObject(getInvoiceURL,"GET");
        Invoice invoice = new Invoice(
                
        );
        /* Check if time of request >12 hrs of booking time */

        /* cancel booking */

        return 2;
    }

    @Override
    public int confirmPayment(int booking_id, long price, String payer_name, int payment_type) {
        return 0;
    }


}
