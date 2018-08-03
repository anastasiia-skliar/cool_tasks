package entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import java.util.Date;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Task {
    private String ID;
    private String uID;
    private String Name;
    private Date Time;
    private Date CreatedAt;
    private Date UpdatedUp;
    private String Desc;

    @Override
    public String toString() {
        return "Task{" +
                "ID='" + ID + '\'' +
                ", uID='" + uID + '\'' +
                ", Name='" + Name + '\'' +
                ", Time=" + Time +
                ", CreatedAt=" + CreatedAt +
                ", UpdatedUp=" + UpdatedUp +
                ", Desc='" + Desc + '\'' +
                '}';
    }

    public String getID() {
        return ID;
    }

    public void setID(String ID) {
        this.ID = ID;
    }

    public String getuID() {
        return uID;
    }

    public void setuID(String uID) {
        this.uID = uID;
    }

    public String getName() {
        return Name;
    }

    public void setName(String name) {
        Name = name;
    }

    public Date getTime() {
        return Time;
    }

    public void setTime(Date time) {
        Time = time;
    }

    public Date getCreatedAt() {
        return CreatedAt;
    }

    public void setCreatedAt(Date createdAt) {
        CreatedAt = createdAt;
    }

    public Date getUpdatedUp() {
        return UpdatedUp;
    }

    public void setUpdatedUp(Date updatedUp) {
        UpdatedUp = updatedUp;
    }

    public String getDesc() {
        return Desc;
    }

    public void setDesc(String desc) {
        Desc = desc;
    }
}
