package com.kliver.domains.api.control;

import android.view.View;
import com.kliver.domains.api.R;
import com.kliver.domains.api.view.MainActivity;

public class MainController implements View.OnClickListener {

    private MainActivity activity;

    public MainController(MainActivity activity) {
        this.activity = activity;

        activity.getBtnDomain().setOnClickListener(this);
        activity.getBtnHistory().setOnClickListener(this);
    }

    @Override
    public void onClick(View v) {
        switch(v.getId()) {
            case R.id.btnDomain:
                activity.openDomainActivity();
                break;
            case R.id.btnHistory:
                activity.openHistoryActivity();
                break;
        }
    }
}
