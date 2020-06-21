package com.kliver.domains.api.view;

import androidx.appcompat.app.AppCompatActivity;
import android.content.Intent;
import android.os.Bundle;
import android.widget.Button;
import com.kliver.domains.api.R;
import com.kliver.domains.api.control.MainController;
import com.kliver.domains.api.util.Constants;

public class MainActivity extends AppCompatActivity {

    private MainController controller;
    private Button btnDomain;
    private Button btnHistory;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        btnDomain = findViewById(R.id.btnDomain);
        btnHistory = findViewById(R.id.btnHistory);

        controller = new MainController(this);
    }

    public void openDomainActivity() {
        Intent i = new Intent(this,DomainActivity.class);
        this.startActivity(i);
    }

    public void openHistoryActivity() {
        Intent i = new Intent(this,HistoryActivity.class);
        this.startActivity(i);
    }

    public Button getBtnDomain() {
        return btnDomain;
    }

    public Button getBtnHistory() {
        return btnHistory;
    }

}
