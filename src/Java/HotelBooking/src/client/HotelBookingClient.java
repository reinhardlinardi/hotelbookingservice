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
        System.out.println("Book Room = "+hb.orderRoom("ReinhardCarry","12345","reinhardcarry@example.com",3,"2018-11-14","2018-11-15"));
        System.out.println("Validate Payment = "+hb.confirmPayment(1,2,"A",3));
    }
}
