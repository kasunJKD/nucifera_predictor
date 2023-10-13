import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit {
  data: any; // Variable to store the response data

  lstmplotFit: SafeResourceUrl;
  lstmPlotValidation: SafeResourceUrl;
  lstmTestPredictions: SafeResourceUrl;

  gruplotFit: SafeResourceUrl;
  gruPlotValidation: SafeResourceUrl;
  gruTestPredictions: SafeResourceUrl;


  dplotFit: SafeResourceUrl;
  dPlotValidation: SafeResourceUrl;
  dTestPredictions: SafeResourceUrl;

  constructor(private http: HttpClient, private sanitizer: DomSanitizer) {}

  ngOnInit() {
    // Send an HTTP POST request to the API endpoint when the component is initialized.
    const batchId = 1; // Set your batch ID here
    const requestData = { batchNumber : batchId };

    this.http.post('http://localhost:11001/api/batch/data/get', requestData).subscribe((response) => {
      this.data = response;
      // const base64Image = this.data.batchResponse[0].plotFit;
      // const decodedImageData = atob(base64Image);
      // const imageUrl = 'data:image/png;base64,' + decodedImageData;

      // this.imageSrc = this.sanitizer.bypassSecurityTrustResourceUrl(imageUrl);

      if (this.data.batchResponse[0])
      {
        this.lstmplotFit = this.sanitizer.bypassSecurityTrustResourceUrl('data:image/png;base64,' + atob(this.data.batchResponse[0].plotFit));
        this.lstmPlotValidation = this.sanitizer.bypassSecurityTrustResourceUrl('data:image/png;base64,' + atob(this.data.batchResponse[0].plotValidation));
        this.lstmTestPredictions = this.sanitizer.bypassSecurityTrustResourceUrl('data:image/png;base64,' + atob(this.data.batchResponse[0].testPredictions));
      }
      if (this.data.batchResponse[1])
      {
        this.gruplotFit = this.sanitizer.bypassSecurityTrustResourceUrl('data:image/png;base64,' + atob(this.data.batchResponse[1].plotFit));
        this.gruPlotValidation = this.sanitizer.bypassSecurityTrustResourceUrl('data:image/png;base64,' + atob(this.data.batchResponse[1].plotValidation));
        this.gruTestPredictions = this.sanitizer.bypassSecurityTrustResourceUrl('data:image/png;base64,' + atob(this.data.batchResponse[1].testPredictions));
      }
      if (this.data.batchResponse[2])
      {
        this.dplotFit = this.sanitizer.bypassSecurityTrustResourceUrl('data:image/png;base64,' + atob(this.data.batchResponse[2].plotFit));
        this.dPlotValidation = this.sanitizer.bypassSecurityTrustResourceUrl('data:image/png;base64,' + atob(this.data.batchResponse[2].plotValidation));
        this.dTestPredictions = this.sanitizer.bypassSecurityTrustResourceUrl('data:image/png;base64,' + atob(this.data.batchResponse[2].testPredictions));
      }
    });
  }
}
