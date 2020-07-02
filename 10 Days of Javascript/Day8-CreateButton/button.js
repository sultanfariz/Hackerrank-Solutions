const btn = document.getElementById('btn');
btn.innerHTML = 0;
btn.addEventListener('click', function () {
    btn.innerHTML = +(btn.innerHTML) + 1;
});