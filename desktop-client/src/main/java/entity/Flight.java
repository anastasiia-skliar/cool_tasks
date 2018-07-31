package entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import java.util.Date;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Flight {
    private String ID;
    private String DepartureCity;
    private Date DepartureTime;
    private String ArrivalCity;
    private Date ArrivalDate;
    private int Price;

    @Override
    public String toString() {
        return "Flight{" +
                "ID='" + ID + '\'' +
                ", DepartureCity='" + DepartureCity + '\'' +
                ", DepartureTime=" + DepartureTime +
                ", ArrivalCity='" + ArrivalCity + '\'' +
                ", ArrivalDate=" + ArrivalDate +
                ", Price=" + Price +
                '}';
    }

    public String getID() {
        return ID;
    }

    public void setID(String ID) {
        this.ID = ID;
    }

    public String getDepartureCity() {
        return DepartureCity;
    }

    public void setDepartureCity(String departureCity) {
        DepartureCity = departureCity;
    }

    public Date getDepartureTime() {
        return DepartureTime;
    }

    public void setDepartureTime(Date departureTime) {
        DepartureTime = departureTime;
    }



    public String getArrivalCity() {
        return ArrivalCity;
    }

    public void setArrivalCity(String arrivalCity) {
        ArrivalCity = arrivalCity;
    }



    public Date getArrivalDate() {
        return ArrivalDate;
    }

    public void setArrivalDate(Date arrivalDate) {
        ArrivalDate = arrivalDate;
    }

    public int getPrice() {
        return Price;
    }

    public void setPrice(int price) {
        Price = price;
    }
}
