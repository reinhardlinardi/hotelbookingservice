package hotelbooking;

import javax.jws.WebService;

@WebService(endpointInterface = "hotelbooking.HotelBookingInterface")
public class HotelBookingImplementation implements HotelBookingInterface{

//    /* Main */
//
//    public static void main(RoomType[] argv) {
//        Object implementor = new HotelBookingImplementation();
//        String address = "http://localhost:9000/HelloWorld";
//        Endpoint.publish(address, implementor);
//    }

    /* Web services */

    @Override
    public int orderRoom(String name, String id_card, String email, int room_type_id, String check_in, String check_out) {
        return 0;
    }

    @Override
    public int cancel(int bookingId) {
        return 2;
    }

    @Override
    public int confirmPayment(int booking_id, long price, String payer_name, int payment_type) {
        return 0;
    }


}
