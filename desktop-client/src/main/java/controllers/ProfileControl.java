package controllers;

import com.jfoenix.controls.JFXTextField;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;

import entity.User;
import java.net.URL;
import java.util.ResourceBundle;

public class ProfileControl extends ControlWithSide implements Initializable {

    @FXML
    JFXTextField id;
    @FXML
    JFXTextField Name;
    @FXML
    JFXTextField Login;


    @Override
    public void initialize(URL location, ResourceBundle resources) {
        initializeSide(location, resources);
        id.setText(LoginControl.userInfo.getID());
        Name.setText(LoginControl.userInfo.getName());
        Login.setText(LoginControl.userInfo.getLogin());

    }
}