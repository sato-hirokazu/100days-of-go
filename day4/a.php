<?php
$client_id = 'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9';
$client_secret = 'eyJ0eXAiOiJK45451NiJ9';
  
 $header = base64_encode($client_id . ':' . $client_secret);
  
 $params = array(
        'grant_type' => 'refresh_token',
        'refresh_token' => 'AHES6ZQPsv3rIA4sAWp3DDZNLXcsLhJc1sN4usryx7BQWWbcPXUdiQ',
 );
  
 // curlでPOST送信
 $ch=curl_init();
 curl_setopt($ch, CURLOPT_URL, 'https://gateway.dmm.com/connect/v1/token');
 curl_setopt($ch, CURLOPT_POST, 1);
 curl_setopt($ch, CURLOPT_HTTPHEADER, array('Authorization: Basic ' . $header));
 curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($params));
 curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE); // 開発環境の場合のみ
 curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
 $res = curl_exec($ch);
 curl_close($ch);

var_dump($res);
?>
