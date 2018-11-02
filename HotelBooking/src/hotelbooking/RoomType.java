package hotelbooking;
public enum RoomType {

    /* Enumerations */
    SINGLE(1), DOUBLE(2), FAMILY(3);

    /* Properties */
    int roomType;

    /* Methods */

    // Constructors
    RoomType() {}

    RoomType(int roomType) { this.roomType = roomType; }
}
