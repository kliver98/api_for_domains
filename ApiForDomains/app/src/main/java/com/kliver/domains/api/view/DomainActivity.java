package com.kliver.domains.api.view;

import androidx.appcompat.app.AppCompatActivity;
import com.kliver.domains.api.R;

import android.os.Bundle;

public class DomainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_domain);
    }

    @Override
    public void onBackPressed() {
        super.onBackPressed();
        this.finish();
    }

}
