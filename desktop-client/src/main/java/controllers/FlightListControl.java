package controllers;

import entity.Flight;
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


public class FlightListControl extends ControlWithSide implements Initializable {

    public static Flight[] flights;

    @FXML
    Button addFlightToTrip;

    @FXML
    private TableView flightTable;
    @FXML
    private TableColumn flightID;
    @FXML
    private TableColumn departureCity;
    @FXML
    private TableColumn departureTime;
    @FXML
    private TableColumn arrivalCity;
    @FXML
    private TableColumn arrivalTime;
    @FXML
    private TableColumn price;

    @Override
    public void initialize(URL location, ResourceBundle resources) {
        initializeSide(location, resources);
        setFlightTable();
    }

    private void setFlightTable() {
        ObservableList flightsList = FXCollections.observableList(Arrays.asList(flights));
        flightTable.setItems(flightsList);
        flightID.setCellValueFactory(new PropertyValueFactory("ID"));
        departureCity.setCellValueFactory(new PropertyValueFactory("DepartureCity"));
        departureTime.setCellValueFactory(new PropertyValueFactory("DepartureTime"));

        arrivalCity.setCellValueFactory(new PropertyValueFactory("ArrivalCity"));
        arrivalTime.setCellValueFactory(new PropertyValueFactory("ArrivalTime"));

        price.setCellValueFactory(new PropertyValueFactory("Price"));
        flightTable.getColumns().setAll(flightID,departureCity,departureTime,arrivalCity,arrivalTime,price);
    }
}
