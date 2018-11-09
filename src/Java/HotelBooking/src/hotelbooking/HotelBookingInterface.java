package hotelbooking;

import javax.jws.WebMethod;
import javax.jws.WebService;

@WebService
public interface HotelBookingInterface {
    @WebMethod int orderRoom(String name, String id_card, String email, int room_type_id, String check_in, String check_out);
    /*
    input:
    name (String)
    id_card (String)
    email(String)
    room_type_id(int)
    check_in (date)
    check_out (date)
    output:
    booking_id (int)
     */
    @WebMethod int cancel(int bookingId);

    @WebMethod int confirmPayment(int booking_id, long price, String payer_name, int payment_type);
    /*
    input :
    booking_id(int)
    price(long)
    payer_name(String)
    payment_type(int)
    output:
    booking_id(int)
     */

}
