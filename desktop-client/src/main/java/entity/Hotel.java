package entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Hotel {
    public String getID() {
        return ID;
    }

    public void setID(String ID) {
        this.ID = ID;
    }

    public String getName() {
        return Name;
    }

    @Override
    public String toString() {
        return "Hotel{" +
                "ID='" + ID + '\'' +
                ", Name='" + Name + '\'' +
                ", Capacity=" + Capacity +
                ", RoomsLeft=" + RoomsLeft +
                ", Floors=" + Floors +
                ", MaxPrice='" + MaxPrice + '\'' +
                ", Address='" + Address + '\'' +
                '}';
    }

    public void setName(String name) {
        Name = name;
    }

    public int getCapacity() {
        return Capacity;
    }

    public void setCapacity(int capacity) {
        Capacity = capacity;
    }

    public int getRoomsLeft() {
        return RoomsLeft;
    }

    public void setRoomsLeft(int roomsLeft) {
        RoomsLeft = roomsLeft;
    }

    public int getFloors() {
        return Floors;
    }

    public void setFloors(int floors) {
        Floors = floors;
    }

    public String getMaxPrice() {
        return MaxPrice;
    }

    public void setMaxPrice(String maxPrice) {
        MaxPrice = maxPrice;
    }

    public String getAddress() {
        return Address;
    }

    public void setAddress(String address) {
        Address = address;
    }

    private String ID;
    private String Name;
    private int Capacity;
    private int RoomsLeft;
    private int Floors;
    private String MaxPrice;
    private String Address;
}
