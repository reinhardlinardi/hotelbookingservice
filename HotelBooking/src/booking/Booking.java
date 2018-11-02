package booking;

import javax.jws.WebMethod;
import javax.jws.WebParam;
import javax.jws.WebService;
import javax.xml.ws.Endpoint;

@WebService()
public class Booking {
    @WebMethod
    public String sayHelloWorldFrom(String from) {
        String result = "Hello, world, from " + from;
        System.out.println(result);
        return result;
    }

    @WebMethod
    public Invoice bookRoom(){

        Invoice invoice = new Invoice("Reinhard","2 November","paid","creditCard","123","456",500000,true);
        return invoice;
    }

    @WebMethod
    public String cancelBooking(String bookingId){
        String status = "Booking with ID "+bookingId+" has been cancelled";
        return status;
    }

    @WebMethod
    public boolean validatePayment(String bookingId){
        boolean status = true;
        System.out.println("Booking with id "+bookingId+" has been cancelled");
        return status;
    }

    public static void main(String[] argv) {
        Object implementor = new Booking();
        String address = "http://localhost:9000/HelloWorld";
        Endpoint.publish(address, implementor);
    }
}
