package entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import java.sql.Time;
import java.util.Date;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Museum {
    private String ID;
    private String Name;
    private String Location;
    private int Price;
    private Date OpenedAt;
    private Date ClosedAt;
    private String MuseumType;
    private String Info;

    @Override
    public String toString() {
        return "Museum{" +
                "ID='" + ID + '\'' +
                ", Name='" + Name + '\'' +
                ", Location='" + Location + '\'' +
                ", Price=" + Price +
                ", OpenedAt=" + OpenedAt +
                ", ClosedAt=" + ClosedAt +
                ", MuseumType='" + MuseumType + '\'' +
                ", Info='" + Info + '\'' +
                '}';
    }

    public String getID() {
        return ID;
    }

    public void setID(String ID) {
        this.ID = ID;
    }

    public String getName() {
        return Name;
    }

    public void setName(String name) {
        Name = name;
    }

    public String getLocation() {
        return Location;
    }

    public void setLocation(String location) {
        Location = location;
    }

    public int getPrice() {
        return Price;
    }

    public void setPrice(int price) {
        Price = price;
    }

    public Date getOpenedAt() {
        return OpenedAt;
    }

    public void setOpenedAt(Time openedAt) {
        OpenedAt = openedAt;
    }

    public Date getClosedAt() {
        return ClosedAt;
    }

    public void setClosedAt(Time closedAt) {
        ClosedAt = closedAt;
    }

    public String getMuseumType() {
        return MuseumType;
    }

    public void setMuseumType(String museumType) {
        MuseumType = museumType;
    }

    public String getInfo() {
        return Info;
    }

    public void setInfo(String info) {
        Info = info;
    }
}
