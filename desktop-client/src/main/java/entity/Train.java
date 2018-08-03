package entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import java.util.Date;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Train {
    @Override
    public String toString() {
        return "Train{" +
                "ID='" + ID + '\'' +
                ", DepartureTime=" + DepartureTime +
                ", ArrivalTime=" + ArrivalTime +
                ", DepartureCity='" + DepartureCity + '\'' +
                ", ArrivalCity='" + ArrivalCity + '\'' +
                ", TrainType='" + TrainType + '\'' +
                ", CarType='" + CarType + '\'' +
                ", Price='" + Price + '\'' +
                '}';
    }

    public String getID() {
        return ID;
    }

    public void setID(String ID) {
        this.ID = ID;
    }

    public Date getDepartureTime() {
        return DepartureTime;
    }

    public void setDepartureTime(Date departureTime) {
        DepartureTime = departureTime;
    }


    public Date getArrivalTime() {
        return ArrivalTime;
    }

    public void setArrivalTime(Date arrivalTime) {
        ArrivalTime = arrivalTime;
    }


    public String getDepartureCity() {
        return DepartureCity;
    }

    public void setDepartureCity(String departureCity) {
        DepartureCity = departureCity;
    }

    public String getArrivalCity() {
        return ArrivalCity;
    }

    public void setArrivalCity(String arrivalCity) {
        ArrivalCity = arrivalCity;
    }

    public String getTrainType() {
        return TrainType;
    }

    public void setTrainType(String trainType) {
        TrainType = trainType;
    }

    public String getCarType() {
        return CarType;
    }

    public void setCarType(String carType) {
        CarType = carType;
    }

    public String getPrice() {
        return Price;
    }

    public void setPrice(String price) {
        Price = price;
    }

    private String ID;
    private Date DepartureTime;
    private Date ArrivalTime;
    private String DepartureCity;
    private String ArrivalCity;
    private String TrainType;
    private String CarType;
    private String Price;
}
