package booking;

import javax.jws.WebMethod;
import javax.jws.WebService;
import javax.xml.ws.Endpoint;

@WebService()
public class HotelBooking {

    @WebMethod
    public Invoice bookRoom(String date){

        Invoice invoice = new Invoice("Reinhard",date,"paid","creditCard","123","456",500000,true);
        return invoice;
    }

    @WebMethod
    public String cancelBooking(String bookingId){
        String status = "HotelBooking with ID "+bookingId+" has been cancelled";
        return status;
    }

    @WebMethod
    public boolean validatePayment(String bookingId){
        boolean status = true;
        System.out.println("HotelBooking with id "+bookingId+" has been cancelled");
        return status;
    }

    public static void main(String[] argv) {
        Object implementor = new HotelBooking();
        String address = "http://localhost:9000/HelloWorld";
        Endpoint.publish(address, implementor);
    }
}
