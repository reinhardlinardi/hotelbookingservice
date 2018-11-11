package client;

import hotelbooking.HotelBooking;

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

//        System.out.println("CancelRoom = "+hb.cancel(6));
        System.out.println("Book Room = "+hb.orderRoom("Reinhard","12345","reinhard@example.com",3,"25-11-2018","26-11-2018"));
//        System.out.println("Validate Payment = "+hb.confirmPayment(5,500000,"A",3));
        System.out.println("Done");
    }
}
