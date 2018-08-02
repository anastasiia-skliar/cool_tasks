package controllers;


import entity.Event;
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

public class EventListControl extends ControlWithSide implements Initializable {

    public static Event[] events;
    @FXML
    Button addEventToTrip;
    @FXML
    private TableView eventTable;
    @FXML
    private TableColumn eventID;
    @FXML
    private TableColumn eventTitle;
    @FXML
    private TableColumn eventCategory;
    @FXML
    private TableColumn eventTown;
    @FXML
    private TableColumn eventDate;
    @FXML
    private TableColumn eventPrice;

    @Override
    public void initialize(URL location, ResourceBundle resources) {
        initializeSide(location, resources);
        setEventsTable();
    }

    private void setEventsTable() {
        ObservableList eventsList = FXCollections.observableList(Arrays.asList(events));
        eventTable.setItems(eventsList);
        eventID.setCellValueFactory(new PropertyValueFactory("ID"));
        eventTitle.setCellValueFactory(new PropertyValueFactory("Title"));
        eventCategory.setCellValueFactory(new PropertyValueFactory("Category"));
        eventTown.setCellValueFactory(new PropertyValueFactory("Town"));
        eventDate.setCellValueFactory(new PropertyValueFactory("Date"));
        eventPrice.setCellValueFactory(new PropertyValueFactory("Price"));
        eventTable.getColumns().setAll(eventID, eventTitle, eventCategory, eventTown, eventDate, eventPrice);
    }
}
