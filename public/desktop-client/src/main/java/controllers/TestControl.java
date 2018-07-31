package controllers;


import com.google.gson.Gson;
import entity.User;
import javafx.fxml.FXML;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.TextField;
import rest.ResponceService;
import rest.UserService;

public class TestControl  {

  @FXML
  Button getButton;
  @FXML
  Label responceLabel;
  @FXML
  Button getButton2;
  @FXML
  TextField login;
  @FXML
          TextField password;

    ResponceService responceService = new ResponceService();
    UserService userService = new UserService();

  @FXML
  public void setLabel(){
    //  responceLabel.setText(String.valueOf(userService.getUserByID("03dc3258-86a7-11e8-9a39-d4bed959082a")[0].getName()));
  }

  @FXML
  public void setLabel2(){
    responceLabel.setText(responceService.method(login.getText(),password.getText()));
  }

}



