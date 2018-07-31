package controllers;

import entity.Hotel;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.scene.control.Button;
import javafx.scene.control.TableColumn;
import javafx.scene.control.TableView;
import javafx.scene.control.cell.PropertyValueFactory;

import java.net.URL;
import java.util.Arrays;
import java.util.ResourceBundle;

public class HotelListControl extends ControlWithSide implements Initializable {

    public static Hotel[] hotels;
    @FXML
    Button addHotelToTrip;
    @FXML
    private TableView hotelTable;
    @FXML
    private TableColumn hotelID;
    @FXML
    private TableColumn hotelName;
    @FXML
    private TableColumn hotelCapacity;
    @FXML
    private TableColumn hotelRoomsLeft;
    @FXML
    private TableColumn hotelFloors;
    @FXML
    private TableColumn hotelMaxPrice;
    @FXML
    private TableColumn hotelAddress;


    @Override
    public void initialize(URL location, ResourceBundle resources) {
        initializeSide(location, resources);
        setHotelsTable();
    }

    private void setHotelsTable() {
        ObservableList hotelsList = FXCollections.observableList(Arrays.asList(hotels));
        hotelTable.setItems(hotelsList);
        hotelID.setCellValueFactory(new PropertyValueFactory("ID"));
        hotelName.setCellValueFactory(new PropertyValueFactory("Name"));
        hotelCapacity.setCellValueFactory(new PropertyValueFactory("Capacity"));
        hotelRoomsLeft.setCellValueFactory(new PropertyValueFactory("RoomsLeft"));
        hotelFloors.setCellValueFactory(new PropertyValueFactory("Floors"));
        hotelMaxPrice.setCellValueFactory(new PropertyValueFactory("MaxPrice"));
        hotelAddress.setCellValueFactory(new PropertyValueFactory("Address"));
        hotelTable.getColumns().setAll(hotelID, hotelName, hotelCapacity, hotelRoomsLeft, hotelFloors,hotelMaxPrice,hotelAddress);
    }
}
