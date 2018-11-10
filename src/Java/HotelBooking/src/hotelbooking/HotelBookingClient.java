package hotelbooking;

import javax.xml.namespace.QName;
import javax.xml.ws.Service;
import java.net.URL;

public class HotelBookingClient {
    public static void main(String[] args) throws Exception{
        URL url = new URL("http://localhost:8080/services/HotelBooking?wsdl");
        QName qname = new QName("http://hotelbooking/","HotelBookingImplementationService");
        Service service = Service.create(url,qname);

        qname = new QName("http://hotelbooking/","HotelBookingImplementationPort");

        HotelBooking hb = service.getPort(qname, HotelBooking.class);

        System.out.println("CancelRoom = "+hb.cancel(0));
        System.out.println("Book Room = "+hb.orderRoom("A","A","A",15,"A","A"));
        System.out.println("Validate Payment = "+hb.confirmPayment(1,2,"A",3));
    }
}
