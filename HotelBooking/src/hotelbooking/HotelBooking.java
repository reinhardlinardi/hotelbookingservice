package hotelbooking;

import javax.jws.WebMethod;
import javax.jws.WebService;
import javax.jws.soap.SOAPBinding;
import javax.xml.ws.Endpoint;

@WebService
@SOAPBinding
public class HotelBooking {

    /* Main */

    public static void main(RoomType[] argv) {
        Object implementor = new HotelBooking();
        String address = "http://localhost:9000/HelloWorld";
        Endpoint.publish(address, implementor);
    }

    /* Web services */

    @WebMethod
    public int book(int bookingId) {
        return 0;
    }

    @WebMethod
    public int cancel(int bookingId) {
        return 0;
    }

    @WebMethod
    public boolean validatePayment(int bookingId, int paymentId) {
        return true;
    }
}
