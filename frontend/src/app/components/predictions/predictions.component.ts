import { Component, OnInit} from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-predictions',
  templateUrl: './predictions.component.html',
  styleUrls: ['./predictions.component.css']
})
export class PredictionsComponent implements OnInit{
  data: any;
  data_gru:any;
  data_d:any;

  constructor(private http: HttpClient) {}

  ngOnInit() {
    
    const modelId_lstm = 1;
    const modelId_gru = 2;
    const modelId_dd = 3;

    const requestData_lstm = {modelId:modelId_lstm};
    const requestData_gru = {modelId:modelId_gru};
    const requestData_dd = {modelId:modelId_dd};

    this.http.post('http://localhost:11001/api/batch/predictions/get', requestData_lstm).subscribe((response) => {
      this.data = response;
    });


    this.http.post('http://localhost:11001/api/batch/predictions/get', requestData_gru).subscribe((response) => {
      this.data_gru = response;
    });


    this.http.post('http://localhost:11001/api/batch/predictions/get', requestData_dd).subscribe((response) => {
      this.data_d = response;
    });
  }

}
