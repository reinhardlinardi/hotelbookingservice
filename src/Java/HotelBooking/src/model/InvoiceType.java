package model;

public enum InvoiceType {

    /* Enumerations */
    BOOKING(1), CANCELLATION(2);

    /* Methods */
    int invoiceType;

    // Constructors
    InvoiceType() {}

    InvoiceType(int invoiceType) { this.invoiceType = invoiceType; }
}
