package hotelbooking;

import javax.xml.ws.Endpoint;

public class HotelBookingPublisher {
    public static void main(String[] args){
        Object implementor = new HotelBookingImplementation();
        String address = "http://localhost:8080/service/HotelBookingServices";
        Endpoint.publish(address, implementor);
    }
}
